# Contributing to TriggerX Contracts

Thank you for your interest in contributing to TriggerX Contracts! This document provides guidelines and instructions for contributing to our smart contract infrastructure for Web3 Automation.

## Code of Conduct

In the interest of fostering an open and welcoming environment, we expect all contributors to be respectful, inclusive, and considerate of others. Smart contract development requires careful collaboration, and we value constructive feedback and professional communication.

## How Can I Contribute?

### Reporting Bugs and Security Vulnerabilities

**CRITICAL:** Smart contract bugs can lead to financial loss. Please follow these guidelines carefully.

#### Before Submitting a Bug Report

1. **Check existing issues** to see if the bug has already been reported
2. **Verify the bug** on the latest version of the contracts
3. **Assess the severity** - is this a security vulnerability or a functional bug?

#### For Security Vulnerabilities

If you discover a security vulnerability or exploit:

1. **DO NOT create a public issue**
2. Contact the maintainers privately through secure channels
3. Wait for acknowledgment before disclosing publicly
4. Follow responsible disclosure practices

For non-critical vulnerabilities that you wish to demonstrate:

1. Create a **Forge test script** that demonstrates the exploit
2. Place it in a separate test file (e.g., `test/exploits/YourExploit.t.sol`)
3. Document the vulnerability clearly in comments
4. Include steps to reproduce and expected vs actual behavior

#### For Functional Bugs

When submitting a bug report, you **MUST** provide one of the following:

1. **Forge Test Script** (Preferred):
   ```solidity
   // test/bugs/BugDemonstration.t.sol
   pragma solidity ^0.8.27;
   
   import "forge-std/Test.sol";
   import "../src/YourContract.sol";
   
   contract BugDemonstration is Test {
       // Setup and test demonstrating the bug
       function testDemonstrateBug() public {
           // Clear reproduction steps
           // Expected behavior vs actual behavior
       }
   }
   ```

2. **Transaction Hash** on a testnet (Sepolia, Holesky, etc.) showing the bug

3. **Foundry Cast Simulation** output demonstrating the issue:
   ```bash
   cast call <contract_address> "functionName()" --rpc-url <network>
   ```

#### Bug Report Template

```markdown
## Bug Description
[Clear description of the bug]

## Severity
- [ ] Critical (funds at risk, contract unusable)
- [ ] High (major functionality broken)
- [ ] Medium (feature bug, workaround exists)
- [ ] Low (minor issue, cosmetic)

## affected Contracts
- Contract Name: [e.g., TaskExecutionHub.sol]
- Function: [e.g., executeTask()]
- Line Numbers: [if applicable]

## Reproduction

### Forge Test Script
[Paste your test script or link to file]

OR

### Transaction Hash
- Network: [e.g., Sepolia]
- Tx Hash: `0x...`
- Contract Address: `0x...`

OR

### Cast Simulation
[Paste command and output]

## Expected Behavior
[What should happen]

## Actual Behavior
[What actually happens]

## Additional Context
- Foundry Version: [Run `forge --version`]
- Solidity Version: [e.g., 0.8.27]
- Any relevant logs or stack traces
```

### Feature Requests

We welcome ideas for new features and improvements!

#### Feature Request Template

```markdown
## Feature Description
[Clear description of the proposed feature]

## Use Case
[Why is this feature valuable? What problem does it solve?]

## Proposed Implementation
[If you have ideas about how to implement this]

## Smart Contract Implications
- [ ] New contract required
- [ ] Modification to existing contract
- [ ] Gas optimization
- [ ] Security considerations

## Examples
[Code snippets, diagrams, or examples of how this would work]

## Alternatives Considered
[Other approaches you've thought about]
```

### Pull Requests

#### Before Submitting a PR

1. **Fork the repository** and create your branch from `main`
2. **Follow the development workflow** outlined below
3. **Write or update tests** for your changes
4. **Run the full test suite** and ensure all tests pass
5. **Run static analysis** (Slither, if applicable)
6. **Update documentation** to reflect your changes
7. **Follow our coding standards**

#### PR Template Guidelines

Your PR should include:

1. **Clear title** describing the change
2. **Description** of what changed and why
3. **Related issues** (e.g., "Fixes #123" or "Relates to #456")
4. **Test coverage** - what tests did you add/modify?
5. **Gas impact** - does this change gas costs? By how much?
6. **Security considerations** - any security implications?
7. **Breaking changes** - does this break existing functionality?

## Development Workflow

### Prerequisites

- [Foundry](https://book.getfoundry.sh/getting-started/installation) installed
- Node.js (for additional tooling)
- Git

### Initial Setup

1. **Fork and clone** the repository:
   ```bash
   git clone https://github.com/YOUR_USERNAME/triggerx-contracts.git
   cd triggerx-contracts/contracts
   ```

2. **Install dependencies**:
   ```bash
   forge install
   ```

3. **Copy environment variables**:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Build contracts**:
   ```bash
   forge build
   ```

### Development Cycle

1. **Create a new branch**:
   ```bash
   git checkout -b feature/your-feature-name
   # or
   git checkout -b fix/issue-description
   ```

2. **Make your changes** to the contracts in `src/`

3. **Write tests** in `test/`:
   ```bash
   # Run specific test
   forge test --match-test testYourFunction
   
   # Run with verbosity to see detailed output
   forge test -vvvv
   
   # Run with gas reporting
   forge test --gas-report
   ```

4. **Check code coverage**:
   ```bash
   forge coverage
   ```

5. **Run static analysis** (if Slither is set up):
   ```bash
   slither .
   ```

6. **Format your code**:
   ```bash
   forge fmt
   ```

7. **Commit your changes**:
   ```bash
   git add .
   git commit -m "feat: add new feature"
   # or
   git commit -m "fix: resolve issue with X"
   ```

8. **Push and create PR**:
   ```bash
   git push origin feature/your-feature-name
   ```

## Coding Standards

### Solidity Style Guide

Follow the [Solidity Style Guide](https://docs.soliditylang.org/en/latest/style-guide.html) and these additional conventions:

1. **Formatting**:
   - Use `forge fmt` to auto-format code
   - 4 spaces for indentation
   - 100 character line length
   - Use double quotes for strings

2. **Naming Conventions**:
   - Contracts: `PascalCase` (e.g., `TaskExecutionHub`)
   - Functions: `camelCase` (e.g., `executeTask`)
   - Private/Internal: prefix with `_` (e.g., `_validateTask`)
   - Constants: `SCREAMING_SNAKE_CASE` (e.g., `MAX_GAS_PRICE`)
   - Events: `PascalCase` (e.g., `TaskExecuted`)

3. **Comments and Documentation**:
   - Use NatSpec for all public/external functions:
     ```solidity
     /// @notice Executes a scheduled task
     /// @param taskId The unique identifier of the task
     /// @return success Whether the execution was successful
     function executeTask(uint256 taskId) external returns (bool success) {
         // Implementation
     }
     ```
   - Explain complex logic with inline comments
   - Document security considerations
   - Add TODO/FIXME comments where appropriate

4. **Security Best Practices**:
   - Follow Checks-Effects-Interactions pattern
   - Use reentrancy guards where appropriate
   - Validate all inputs
   - Use SafeMath operations (or Solidity 0.8+ overflow protection)
   - Be careful with delegatecall and external calls
   - Document any assembly usage thoroughly

5. **Gas Optimization**:
   - Use `calldata` instead of `memory` for read-only function parameters
   - Pack struct variables efficiently
   - Use `uint256` instead of smaller uints (unless packing)
   - Cache storage variables in memory when accessed multiple times
   - Use events for data that doesn't need to be on-chain

## Gas Optimization Guidelines

When optimizing for gas:

1. **Measure first**: Use `forge snapshot` to establish baseline
2. **Optimize hot paths**: Focus on frequently called functions
3. **Don't sacrifice readability**: Document complex optimizations
4. **Test thoroughly**: Ensure optimizations don't introduce bugs
5. **Document trade-offs**: Note any security/readability trade-offs

## Getting Help

- **Questions**: Open a discussion on GitHub
- **Bugs**: Follow the bug report template
- **Security**: Contact maintainers privately
- **General**: Join our community channels

## License

By contributing to TriggerX Contracts, you agree that your contributions will be licensed under the same license as the project.

---

**Thank you for contributing to TriggerX Contracts! Your efforts help build a more secure and efficient Web3 automation infrastructure.**
