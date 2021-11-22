- [Practicals](#practicals)
  - [Remove useless files](#remove-useless-files)

# Practicals

- [TP2](tp2/README.md)
- [TP3](tp3/README.md)

## Remove useless files

Can be useful because all files are copied in the initial.zip file of the professor.

**Bash**

```bash
rm -rf .gradle .classpath .project build .settings
./gradlew eclipse
```

**Powershell**

```powershell
rm .gradle .classpath .project build .settings
.\gradlew eclipse
```
