const { ethers } = require("hardhat");

async function main() {
    // Fetch the list of accounts
    const [sender, receiver] = await ethers.getSigners();

    // Amount of ETH to send
    const amount = ethers.parseEther("1.0");  // 1 ETH

    // Sending ETH
    const tx = await sender.sendTransaction({
        to: "0x13a97F1E5F1cC0f7A519b5bd9cffefb947e7d589",
        value: amount
    });
    await tx.wait().then((receipt) => {
        console.log(receipt);  // This logs the receipt to see if it was successful
    }).catch((error) => {
        console.error("Transaction failed:", error);
    });

    console.log(sender.balance)
    console.log(`Sent ${ethers.formatEther(amount)} ETH from ${sender.address} to 0x13a97F1E5F1cC0f7A519b5bd9cffefb947e7d589`);
    const balance = await ethers.provider.getBalance("0x13a97F1E5F1cC0f7A519b5bd9cffefb947e7d589");
    console.log(`Balance: ${ethers.formatEther(balance)} ETH`);
}

main().catch((error) => {
    console.error(error);
    process.exit(1);
});