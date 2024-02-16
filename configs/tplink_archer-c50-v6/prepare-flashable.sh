#!/usr/bin/env bash
# according to the instructions you need to download a stock firmware file and extract its uboot
# bootloader manually to prepend it to the sysupgrade image when flashing via the stock firmware.
# OTHERWISE YOU WILL BRICK YOUR ROUTER.
# https://github.com/openwrt/openwrt/pull/13547/commits/e9ac1b19e0f3d383ab83373c261bfb5527a29521

# download a recent firmware archive, see https://www.tp-link.com/de/support/download/archer-c50/#Firmware
firmware="https://static.tp-link.com/upload/firmware/2023/202309/20230925/Archer%20C50(EU)_V6_230810.zip"
echo "download a stock firmware $(basename "$firmware") ..."
curl -LO -# "$firmware"

# extract the bootloader from firmware file within
echo "extract bootloader from binary file within ..."
unzip -p "$(basename "$firmware")" "Archer_C50v6_*.bin" | head -c 131584 > bootloader.bin

# then prepend it to all the compiled sysupgrades
echo "prepend the bootloader to all built sysupgrades ..."
for upgrade in **/openwrt-*-squashfs-sysupgrade.bin; do
  out="${upgrade%-sysupgrade.bin}-stockflashable.bin"
  echo "+ $out"
  cat bootloader.bin "$upgrade" > "$out"
done
