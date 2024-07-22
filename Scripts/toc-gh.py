import json
import shutil
import subprocess
import typing
import re

toc_start_comment = '<!--toc start-->'
toc_end_comment = '<!--toc end-->'
placeholder = fr'{toc_start_comment}[\s\S]+{toc_end_comment}'
owner = 'dhay3'
repo = 'memos'
base_path = 'Docs/'


def transverse_dir2toc(owner: str, repo: str, path: str, toc: typing.List) -> typing.List:
    if shutil.which("gh"):
        cmd = ['gh', 'api',
               '--paginate',
               '-H', 'Accept: application/vnd.github+json',
               '-H', 'X-GitHub-Api-Version: 2022-11-28',
               f'/repos/{owner}/{repo}/contents/{path}']
        proc = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
        stdout, _ = proc.communicate()
        data = json.loads(stdout)
        for entry in data:
            # print(json.dumps(entry, indent=2))
            if 'dir' == entry['type']:
                path = entry['path']
                name = path.split('/')[-1]
                if not name.startswith('.'):
                    md_indent = '  ' * path.replace(base_path, '').count('/')
                    md_link = f'[{name}]({path})'
                    md_list = f'{md_indent}- {md_link}'
                    print(md_list)
                    toc.append(md_list)
                    transverse_dir2toc(owner, repo, path, toc)
    return toc


def toc2readme():
    import uuid
    seed = uuid.uuid4()
    toc = '\n'.join(transverse_dir2toc(owner, repo, base_path, []))
    toc_content = f"""
{toc_start_comment}
<!--{seed}-->
{toc}
{toc_end_comment} 
"""
    with open('../README.md', 'r') as f:
        readme_content = f.read()
    with open('../README.md', 'w') as f:
        n_readme_content = re.sub(placeholder, toc_content, readme_content)
        f.write(n_readme_content)


if __name__ == '__main__':
    toc2readme()
