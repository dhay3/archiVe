---
title: Conda 02 - CLI
author: "0x00"
createTime: 2024-05-30
lastModifiedTime: 2024-05-30-09:03
draft: true
tags:
  - Python
  - AI
---

Conda 大体上的命令和 `pip` 类似，可以参考对比下表[^1]

|Task|Conda package and environment manager command|Pip package manager command|Virtualenv environment manager command|
|---|---|---|---|
|Install a package|`conda install $PACKAGE_NAME`|`pip install $PACKAGE_NAME`|X|
|Update a package|`conda update --name $ENVIRONMENT_NAME $PACKAGE_NAME`|`pip install --upgrade $PACKAGE_NAME`|X|
|Update package manager|`conda update conda`|Linux/macOS: `pip install -U pip` Win: `python -m pip install -U pip`|X|
|Uninstall a package|`conda remove --name $ENVIRONMENT_NAME $PACKAGE_NAME`|`pip uninstall $PACKAGE_NAME`|X|
|Create an environment|`conda create --name $ENVIRONMENT_NAME python`|X|`cd $ENV_BASE_DIR; virtualenv $ENVIRONMENT_NAME`|
|Activate an environment|`conda activate $ENVIRONMENT_NAME`*|X|`source $ENV_BASE_DIR/$ENVIRONMENT_NAME/bin/activate`|
|Deactivate an environment|`conda deactivate`|X|`deactivate`|
|Search available packages|`conda search $SEARCH_TERM`|`pip search $SEARCH_TERM`|X|
|Install package from specific source|`conda install --channel $URL $PACKAGE_NAME`|`pip install --index-url $URL $PACKAGE_NAME`|X|
|List installed packages|`conda list --name $ENVIRONMENT_NAME`|`pip list`|X|
|Create requirements file|`conda list --export`|`pip freeze`|X|
|List all environments|`conda info --envs`|X|Install virtualenv wrapper, then `lsvirtualenv`|
|Install other package manager|`conda install pip`|`pip install conda`|X|
|Install Python|`conda install python=x.x`|X|X|
|Update Python|`conda update python`*|X|X|

也可以参考 Cheatsheet[^2]

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]: [Commands — conda 24.5.0 documentation](https://docs.conda.io/projects/conda/en/stable/commands/index.html)
[^2]: [Cheat sheet — conda 24.5.0 documentation](https://docs.conda.io/projects/conda/en/stable/user-guide/cheatsheet.html)

