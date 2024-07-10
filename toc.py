import os
import subprocess

if __name__ == '__main__':
    subprocess.Popen('pwd')
    base_path = 'Docs'
    for root, dirs, files in os.walk(base_path):
        for dir in dirs:
            print(dir)
            relative_path = os.path.join(root, dir).replace('', '')
            relative_name = relative_path.split(os.sep)[-1]
            indent = '  ' * relative_path.replace(base_path, '').count(os.sep)
            md_link = f'[{relative_name}]{relative_path}'
            md_list = f'{indent}- {md_link}'
            print(md_list)
