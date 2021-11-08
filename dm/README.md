- [Data mining for big data](#data-mining-for-big-data)
  - [Windows (Windows)](#windows-windows)
  - [Unix (Bash)](#unix-bash)

# Data mining for big data

Create env and install packages

## Windows (Windows)

```powershell
python -m venv .venv
.venv/Scripts/activate.ps1
python -m pip -r requirements.txt
```

Maybe this command can help


```powershell
$env:PYSPARK_PYTHON="python"
```


## Unix (Bash)

```bash
python -m venv .venv
source .venv/bin/activate
python -m pip -r requirements.txt
```
