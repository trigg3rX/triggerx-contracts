const port = process.env.HOST_PORT || 9090

module.exports = {
  networks: {
    mainnet: {
      privateKey: process.env.PRIVATE_KEY_MAINNET,
      userFeePercentage: 100,
      feeLimit: 1000 * 1e6,
      fullHost: 'https://api.trongrid.io',
      network_id: '1'
    },
    shasta: {
      privateKey: process.env.SHASTA_PRIVATE_KEY,
      userFeePercentage: 100,
      feeLimit: 1000 * 1e6,
      fullHost: 'https://api.shasta.trongrid.io',
      network_id: '1000'
    },
  
    compilers: {
      solc: {
        version: '0.8.20'
      }
    }
  },
  // solc compiler optimize
  solc: {
  //   optimizer: {
  //     enabled: true,
  //     runs: 200
  //   },
  //   evmVersion: 'istanbul'
  }
}
