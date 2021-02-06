# windows删除右键新建多余选项

HKEY_CLASSES_ROOT下找来到右键自新建菜单子百项度,删除其下的知ShellNew,例如道:
HKEY_CLASSES_ROOT\.bmp\ShellNew
HKEY_CLASSES_ROOT\.contact\ShellNew
HKEY_CLASSES_ROOT\.jnt\jntfile\ShellNew
HKEY_CLASSES_ROOT\.lnk\ShellNew
HKEY_CLASSES_ROOT\.rtf\ShellNew
HKEY_CLASSES_ROOT\.txt\ShellNew
HKEY_CLASSES_ROOT\.zip\CompressedFolder\ShellNew
HKEY_CLASSES_ROOT\Briefcase\ShellNew
HKEY_CLASSES_ROOT\Folder\ShellNew

https://zhidao.baidu.com/question/331448109.html