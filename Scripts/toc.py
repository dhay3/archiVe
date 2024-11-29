import os
import re
import uuid
import subprocess
import urllib.parse as parse
from typing import List, Callable

README_PATH = "README.md"


def transverse_dir2toc(path: str, level: int = 0, flag: bool = True) -> str:
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
                toc += md_toc + "\n" + transverse_dir2toc(sub_dir, level + 1, False)
    return toc


def transverse_file2toc(path: str) -> str:
    toc = ""
    for root, dirs, files in os.walk(path):
        for file in sorted(files):
            cmd = ["git", "check-ignore", os.path.join(root, file)]
            procs = subprocess.Popen(cmd, stdout=subprocess.DEVNULL)
            try:
                procs.communicate(timeout=30)
            except subprocess.TimeoutExpired:
                procs.kill()
                procs.communicate()
            if 0 != procs.returncode:
                md_name = os.path.basename(file).split('.')[0]
                md_path = parse.quote(os.path.join(root, file))
                md_link = f"[{md_name}]({md_path})"
                md_toc = f"- {md_link}"
                toc += md_toc + "\n"
    return toc


def to_toc_content(base_paths: List[str], toc_start_comment: str, toc_end_comment: str, func: Callable):
    entropy = uuid.uuid4()
    placeholder = rf"{toc_start_comment}[\s\S]+{toc_end_comment}"
    toc = list(map(lambda _: func(_), base_paths))
    toc_content = f"""{toc_start_comment}
<!--{entropy}-->
{"".join(toc)}
{toc_end_comment} """
    import logging
    logging.basicConfig(level=logging.DEBUG)
    logging.info(toc_content)
    with open(README_PATH, "r") as f:
        readme_content = f.read()
    with open(README_PATH, "w") as f:
        n_readme_content = re.sub(placeholder, toc_content, readme_content)
        f.write(n_readme_content)


def toc2readme() -> None:
    to_toc_content(["Notes/"],
                   "<!--notes toc start-->",
                   "<!--notes toc end-->",
                   transverse_dir2toc)

    to_toc_content(["Howto/"],
                   "<!--howto toc start-->",
                   "<!--howto toc end-->",
                   transverse_file2toc)


if __name__ == "__main__":
    toc2readme()
