Set WinScriptHost = CreateObject("WScript.Shell")
WinScriptHost.Run Chr(34) & "main.exe" & Chr(34), 0
Set WinScriptHost = Nothing