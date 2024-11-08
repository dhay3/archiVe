import os
import re
import subprocess
import urllib.parse as parse

toc_start_comment = "<!--toc start-->"
toc_end_comment = "<!--toc end-->"
placeholder = rf"{toc_start_comment}[\s\S]+{toc_end_comment}"
base_path = "Notes/"
readme_path = "README.md"


def transverse_toc(path: str, level: int, flag: bool = True) -> str:
    toc = ""
    with os.scandir(path) as objs:
        sorted_sub_dirs = sorted(
            [
                _.path
                for _ in objs
                if not _.name.startswith(".") and not _.is_symlink() and _.is_dir()
            ]
        )
        for sub_dir in sorted_sub_dirs:
            cmd = ["git", "check-ignore", sub_dir]
            procs = subprocess.Popen(cmd, stdout=subprocess.DEVNULL)
            try:
                procs.communicate(timeout=30)
            except subprocess.TimeoutExpired:
                procs.kill()
                procs.communicate()
            if 0 != procs.returncode:
                md_name = os.path.basename(sub_dir)
                md_path = parse.quote(sub_dir)
                md_link = f"[{md_name}]({md_path})"
                md_toc = f"## {md_link}" if flag else f'{"  " * level}- {md_link}'
                toc += md_toc + "\n" + transverse_toc(sub_dir, level + 1, False)
    return toc


def toc2readme(path: str) -> None:
    import uuid

    seed = uuid.uuid4()
    toc = transverse_toc(base_path, 0)
    toc_content = f"""{toc_start_comment}
<!--{seed}-->
> [!note]
> Created by Github Workflows

{toc}
{toc_end_comment} """
    with open(path, "r") as f:
        readme_content = f.read()
    with open(path, "w") as f:
        n_readme_content = re.sub(placeholder, toc_content, readme_content)
        f.write(n_readme_content)


if __name__ == "__main__":
    # print(transverse_toc(base_path, 0))
    toc2readme(readme_path)
