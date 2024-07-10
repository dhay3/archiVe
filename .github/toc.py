import json
import shutil
import subprocess

toc_start_comment = '<!--toc start-->'
toc_end_comment = '<!--toc end-->'


def transverse_dir(owner, repo, path):
    if shutil.which("gh"):
        cmd = ['gh', 'api',
               '--paginate',
               '-H', 'Accept: application/vnd.github+json',
               '-H', 'X-GitHub-Api-Version: 2022-11-28',
               f'/repos/{owner}/{repo}/contents/{path}']
        proc = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
        try:
            stdout, stderr = proc.communicate(timeout=15)
        except subprocess.TimeoutExpired:
            proc.kill()
            stdout, stderr = proc.communicate(timeout=15)
        data = json.loads(stdout)
        for entry in data:
            # print(json.dumps(entry, indent=2))
            type = entry['type']
            if 'dir' == type:
                path = entry['path']
                name = path.split('/')[-1]
                md_link = f'[{name}]({path})'
                print(md_link)
                transverse_dir(owner, repo, path)



if __name__ == '__main__':
    transverse_dir('dhay3', 'memos', 'Docs')
