# This command should be run on a Windows workstation on corp.

# Import all the modules we need.
Get-ChildItem ($PSScriptRoot + "\..\Modules") | foreach {$_.Name} {
    Import-Module -FullyQualifiedName ($PSScriptRoot + "\..\Modules\" + $_)
}

Import-Module PSDesiredStateConfiguration

configuration ChromeEnterpriseLab {

    Node $WindowsNodes.Where({$_.Role == "AD"}).NodeName {
        
        # Need to make sure the host has the correct certificate installed for decrypting secure MOF data.

    }

}
