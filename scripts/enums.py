import clang.cindex
import jinja2

template = jinja2.Template("""// Code generated by scripts/enums.py from "Zydis.h". DO NOT EDIT.
package zydis

{% for name, children in enum_list.items() -%}
type {{name}} int32

const (
{%- for child in children %}
    {{child.name}} {{name}} = {{child.value}}
{%- endfor %}
)

{% endfor %}
""")

def extract_enum_list(file_path):
    clang.cindex.Config.set_library_file('/usr/lib/libclang.so')

    translation_unit = clang.cindex.Index.create().parse(file_path)

    def enum_list(node):
        enums = {}

        if node.kind == clang.cindex.CursorKind.ENUM_DECL:
            name = node.spelling.removeprefix('Zydis').removesuffix('_')

            children = []
            for child in node.get_children():
                if child.kind == clang.cindex.CursorKind.ENUM_CONSTANT_DECL:
                    children.append({'name': child.spelling.removeprefix('ZYDIS_'), 'value': child.enum_value})

            enums[name] = children

        for child in node.get_children():
            enums.update(enum_list(child))

        return enums

    return enum_list(translation_unit.cursor)

with open("zydis_enums.go", "w") as f:
    f.write(template.render(enum_list = extract_enum_list('/usr/include/Zydis/Zydis.h')))
