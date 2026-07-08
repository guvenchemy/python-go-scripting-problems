#!/bin/bash

shell_users=()
rogue_root_users=()
total_users=0

while IFS=":" read -r user pass uid gid desc homedir shell ; do
    echo "user: $user, uid: $uid, gid: $gid, shell: $shell"
    if [[ "$uid" -eq 0 && "$user" != "root" ]]; then
        rogue_root_users+=("$user")
    fi
    if [[ "$shell" != "/sbin/nologin" && "$shell" != "/bin/false" ]]; then
        shell_users+=("$user")
    fi 
    total_users=$((total_users+1))

done < "mock_passwd.txt"

echo "Total number of terminal users in the system: ${#shell_users[@]}"
if [[ "${#rogue_root_users[@]}" -gt 0 ]]; then
    echo "Number of rogue root users: ${#rogue_root_users[@]}"
    echo "[CRITICAL ALERT]: UNAUTHORIZED ROOT ACCESS DETECTED!: ${rogue_root_users[@]}"
fi
echo "Total users: $total_users"
