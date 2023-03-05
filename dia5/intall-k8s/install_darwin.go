package main

func GetKubeCTLInstallCommand() (string, []string) {
	cmd := "brew"
	args := []string{"instal", "kubectl..."}

	return cmd, args
}
