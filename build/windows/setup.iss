[Setup]
AppName=gitgeist
AppVersion=1.0.0
DefaultDirName={pf}\gitgeist
DefaultGroupName=gitgeist
OutputDir=.
OutputBaseFilename=gitgeistinstaller
Compression=lzma
SolidCompression=yes
DisableProgramGroupPage=yes
WizardStyle=modern

[Files]
Source: "gitgeist.exe"; DestDir: "{app}"; Flags: ignoreversion

[Registry]
; Добавляем путь к установленной папке в системный PATH
Root: HKLM; Subkey: "SYSTEM\CurrentControlSet\Control\Session Manager\Environment"; \
    ValueType: expandsz; ValueName: "Path"; ValueData: "{olddata};{app}"; Flags: preservestringtype

[Run]
; Обновляем переменные окружения сразу после установки
Filename: "setx"; Parameters: "PATH ""{app};%PATH%"""; Flags: runhidden

[UninstallDelete]
; Убираем из PATH при удалении
Type: filesandordirs; Name: "{app}"