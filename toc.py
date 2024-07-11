import os
import re
import subprocess
from glob import glob

toc_start_comment = '<!--toc start-->'
toc_end_comment = '<!--toc end-->'
placeholder = fr'{toc_start_comment}[\s\S]+{toc_end_comment}'
base_path = 'Docs/'
readme_path = 'README.md'


def transverse_toc(path: str, indent: str = '') -> str:
    toc = ''
    with os.scandir(path) as objs:
        sorted_sub_dirs = sorted([_.path for _ in objs
                                  if not _.name.startswith('.') and not _.is_symlink() and _.is_dir()])
        for sub_dir in sorted_sub_dirs:
            cmd = ['git', 'check-ignore', sub_dir]
            procs = subprocess.Popen(cmd, stdout=subprocess.DEVNULL)
            try:
                procs.communicate(timeout=30)
            except subprocess.TimeoutExpired:
                procs.kill()
                procs.communicate()
            if 0 != procs.returncode:
                md_name = os.path.basename(sub_dir)
                md_link = f'[{md_name}]({sub_dir})'
                md_list = f'{indent}- {md_link}'
                toc += md_list + '\n'
                toc += transverse_toc(sub_dir, indent + '\t')
    return toc


def toc2readme(path: str) -> None:
    import uuid
    seed = uuid.uuid4()
    toc = transverse_toc(base_path)
    toc_content = f"""
{toc_start_comment}
<!--{seed}-->
{toc}
{toc_end_comment} """
    with open(path, 'r') as f:
        readme_content = f.read()
    with open(path, 'w') as f:
        n_readme_content = re.sub(placeholder, toc_content, readme_content)
        f.write(n_readme_content)


if __name__ == '__main__':
    toc2readme(readme_path)
