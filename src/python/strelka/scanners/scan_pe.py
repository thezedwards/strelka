import binascii
from datetime import datetime
import hashlib
import struct

import pefile

from strelka import strelka


IMAGE_MAGIC_LOOKUP = {
    0x10b: '32_BIT',
    0x20b: '64_BIT',
    0x107: 'ROM_IMAGE',
}


class ScanPe(strelka.Scanner):
    """Collects metadata from PE files."""
    def scan(self, data, file, options, expire_at):
        self.event['total'] = {'sections': 0}

        try:
            pe = pefile.PE(data=data)
            pe_dict = pe.dump_dict()

            self.event['total']['sections'] = pe.FILE_HEADER.NumberOfSections
            self.event['warnings'] = pe.get_warnings()
            self.event['timestamp'] = datetime.utcfromtimestamp(pe.FILE_HEADER.TimeDateStamp).isoformat()
            machine = pe.FILE_HEADER.Machine
            self.event['machine'] = {
                'id': machine,
                'type': pefile.MACHINE_TYPE.get(machine),
            }
            # Reference: http://msdn.microsoft.com/en-us/library/windows/desktop/ms680339%28v=vs.85%29.aspx
            self.event['image_magic'] = IMAGE_MAGIC_LOOKUP.get(pe.OPTIONAL_HEADER.Magic, 'Unknown')
            subsystem = pe.OPTIONAL_HEADER.Subsystem
            self.event['subsystem'] = pefile.SUBSYSTEM_TYPE.get(subsystem)
            self.event['stack_reserve_size'] = pe.OPTIONAL_HEADER.SizeOfStackReserve
            self.event['stack_commit_size'] = pe.OPTIONAL_HEADER.SizeOfStackCommit
            self.event['heap_reserve_size'] = pe.OPTIONAL_HEADER.SizeOfHeapReserve
            self.event['heap_commit_size'] = pe.OPTIONAL_HEADER.SizeOfHeapCommit
            self.event['image_base'] = pe.OPTIONAL_HEADER.ImageBase
            self.event['entry_point'] = pe.OPTIONAL_HEADER.AddressOfEntryPoint
            self.event['image_characteristics'] = pe_dict.get('Flags')
            self.event['dll_characteristics'] = pe_dict.get('DllCharacteristics')

            try:
                self.event['imphash'] = pe.get_imphash()

            except AttributeError:
                self.flags.append('no_imphash')

            self.event.setdefault('export_functions', [])
            export_symbols = pe_dict.get('Exported symbols', [])
            for symbols in export_symbols:
                name = symbols.get('Name')
                if name is not None and isinstance(name, bytes) and name not in self.event['export_functions']:
                    self.event['export_functions'].append(name)

            import_cache = {}
            self.event.setdefault('imports', [])
            import_symbols = pe_dict.get('Imported symbols', [])
            for symbol in import_symbols:
                for import_ in symbol:
                    dll = import_.get('DLL')
                    if dll is not None:
                        if dll not in self.event['imports']:
                            self.event['imports'].append(dll)
                            import_cache.setdefault(dll, [])
                        ordinal = import_.get('Ordinal')
                        if ordinal is not None:
                            ordinal = pefile.ordlookup.ordLookup(dll.lower(), ordinal, make_name=True)
                            import_cache[dll].append(ordinal)
                        name = import_.get('Name')
                        if name is not None:
                            import_cache[dll].append(name)

            self.event.setdefault('import_functions', [])
            for (import_, functions) in import_cache.items():
                import_entry = {'import': import_, 'functions': functions}
                if import_entry not in self.event['import_functions']:
                    self.event['import_functions'].append(import_entry)

            self.event.setdefault('resources', [])
            try:
                for resource in pe.DIRECTORY_ENTRY_RESOURCE.entries:
                    res_type = pefile.RESOURCE_TYPE.get(resource.id, 'Unknown')
                    for entry in resource.directory.entries:
                        for e_entry in entry.directory.entries:
                            sublang = pefile.get_sublang_name_for_lang(
                                e_entry.data.lang,
                                e_entry.data.sublang,
                            )
                            offset = e_entry.data.struct.OffsetToData
                            size = e_entry.data.struct.Size
                            r_data = pe.get_data(offset, size)
                            language = pefile.LANG.get(e_entry.data.lang, 'Unknown')
                            data = {
                                'type': res_type,
                                'id': e_entry.id,
                                'name': e_entry.data.struct.name,
                                'offset': offset,
                                'size': size,
                                'sha256': hashlib.sha256(r_data).hexdigest(),
                                'sha1': hashlib.sha1(r_data).hexdigest(),
                                'md5': hashlib.md5(r_data).hexdigest(),
                                'language': language,
                                'sub_language': sublang,
                            }
                            if data not in self.event['resources']:
                                self.event['resources'].append(data)

            except AttributeError:
                self.flags.append('no_resources')

            if hasattr(pe, 'DIRECTORY_ENTRY_DEBUG'):
                debug = dict()
                for e in pe.DIRECTORY_ENTRY_DEBUG:
                    rawData = pe.get_data(e.struct.AddressOfRawData, e.struct.SizeOfData)
                    if rawData.find(b'RSDS') != -1 and len(rawData) > 24:
                        pdb = rawData[rawData.find(b'RSDS'):]
                        debug['guid'] = b'%s-%s-%s-%s' % (
                            binascii.hexlify(pdb[4:8]),
                            binascii.hexlify(pdb[8:10]),
                            binascii.hexlify(pdb[10:12]),
                            binascii.hexlify(pdb[12:20]),
                        )
                        debug['age'] = struct.unpack('<L', pdb[20:24])[0]
                        debug['pdb'] = pdb[24:].rstrip(b'\x00')
                        self.event['rsds'] = debug
                    elif rawData.find(b'NB10') != -1 and len(rawData) > 16:
                        pdb = rawData[rawData.find(b'NB10') + 8:]
                        debug['created'] = struct.unpack('<L', pdb[0:4])[0]
                        debug['age'] = struct.unpack('<L', pdb[4:8])[0]
                        debug['pdb'] = pdb[8:].rstrip(b'\x00')
                        self.event['nb10'] = debug

            self.event.setdefault('sections', [])
            sections = pe_dict.get('PE Sections', [])
            for section in sections:
                section_entry = {
                    'name': section.get('Name', {}).get('Value', '').replace('\\x00', ''),
                    'flags': section.get('Flags', []),
                    'structure': section.get('Structure', ''),
                }
                if section_entry not in self.event['sections']:
                    self.event['sections'].append(section_entry)

            security = pe.OPTIONAL_HEADER.DATA_DIRECTORY[pefile.DIRECTORY_ENTRY['IMAGE_DIRECTORY_ENTRY_SECURITY']]
            digital_signature_virtual_address = security.VirtualAddress
            if security.Size > 0:
                extract_data = pe.write()[digital_signature_virtual_address + 8:]
                if len(extract_data) > 0:
                    self.flags.append('signed')

                    extract_file = strelka.File(
                        name='digital_signature',
                        source=self.name,
                    )
                    for c in strelka.chunk_string(extract_data):
                        self.upload_to_cache(
                            extract_file.pointer,
                            c,
                            expire_at,
                        )
                    self.files.append(extract_file)

                else:
                    self.flags.append('empty_signature')

            if hasattr(pe, 'FileInfo'):
                self.event.setdefault('version_info', [])
                for structure in pe.FileInfo:
                    for fileinfo in structure:
                        if fileinfo.Key.decode() == 'StringFileInfo':
                            for block in fileinfo.StringTable:
                                for name, value in block.entries.items():
                                    fixedinfo = {
                                        'name': name.decode(),
                                        'value': value.decode(),
                                    }
                                    if fixedinfo not in self.event['version_info']:
                                        self.event['version_info'].append(fixedinfo)
            else:
                self.flags.append('no_version_info')

        except IndexError:
            self.flags.append('index_error')
        except pefile.PEFormatError:
            self.flags.append('pe_format_error')
