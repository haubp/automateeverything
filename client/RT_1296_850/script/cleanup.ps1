New-Item -Path '.\AgentLog' -ItemType Directory

Copy-Item -Path 'C:\ProgramData\OPSWAT' -Destination '.\AgentLog' -recurse -Force

Remove-Item 'C:\ProgramData\OPSWAT' -Force -Recurse
echo "Finish"