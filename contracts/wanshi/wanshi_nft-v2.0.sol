// SPDX-License-Identifier: MIT

pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/access/Ownable.sol";


contract CryptoAssets is Ownable{

    string public standard = 'RoiCryptoAssets';
    string public name;
    string public symbol;
    uint256 public maxSupply;
    string public uriPrefix = "";
    string public uriSuffix = ".json";

    uint256 public faucet = 0.01 ether;

    uint256 public cryptoRemainingToAssign = 0;

    //mapping (address => uint) public addressToAssetsIndex;
    mapping (uint => address) public assetsIndexToAddress;

    /* This creates an array with all balances */
    mapping (address => uint256) public balanceOf;

    event Assign(address indexed to, uint256 assetsIndex);
    event Transfer(address indexed from, address indexed to, uint256 assetsIndex);

    /* Initializes contract with initial supply tokens to the creator of the contract */
    constructor() payable {
        // balanceOf[msg.sender] = initialSupply;              // Give the creator all initial tokens
        maxSupply = 10000;                        // Update total supply
        name = "ROISUPE";                               // Set the name for display purposes
        symbol = "ROIS";                               // Set the symbol for display purposes
    }

    function setInitialOwners(address[] memory addresses, uint256[] memory indices) public {
        uint n = addresses.length;
        for (uint i = 0; i < n; i++) {
            _mintForAddress(addresses[i], indices[i]);
        }
    }

    modifier mintCompliance(uint256 assetsIndex) {
        require(assetsIndex < 10000);
        require(cryptoRemainingToAssign < maxSupply);
        require(assetsIndexToAddress[assetsIndex] == address(0));
        _;
    }

    function _mintForAddress(address to, uint256 assetsIndex) internal{

        assetsIndexToAddress[assetsIndex] = to;
        balanceOf[to] += 1;
        cryptoRemainingToAssign += 1;
        (bool f, ) = payable(to).call{value: faucet}("");
        require(f);
        emit Assign(to, assetsIndex);
    }


    function mint(address to, uint256 assetsIndex) mintCompliance(assetsIndex) public onlyOwner{
        require(assetsIndexToAddress[assetsIndex] != to);

        _mintForAddress(to, assetsIndex);
    }

    function totalSupply() public view returns (uint256) {
        return cryptoRemainingToAssign;
    }

    function addressByIndex(uint256 index) public view returns (address) {
        return assetsIndexToAddress[index];
    }


    function setUriPrefix(string memory _uriPrefix) public onlyOwner{
        uriPrefix = _uriPrefix;
    }

    function setUriSuffix(string memory _uriSuffix) public onlyOwner{
        uriSuffix = _uriSuffix;
    }

    // Transfer ownership of a assets to another user without requiring payment
    function transferNft(address to, uint assetsIndex) public{

        require(assetsIndexToAddress[assetsIndex] == msg.sender);
        require(assetsIndex < 10000);

        assetsIndexToAddress[assetsIndex] = to;
        balanceOf[msg.sender] -= 1;
        balanceOf[to] += 1;
        (bool f, ) = payable(to).call{value: faucet}("");
        require(f);

        emit Transfer(msg.sender, to, 1);
    }

    // Transfer ownership of a assets to another user without requiring payment
    function transferNftFrom(address from, address to, uint assetsIndex) public onlyOwner{

        require(assetsIndexToAddress[assetsIndex] == from);
        require(assetsIndex < 10000);

        assetsIndexToAddress[assetsIndex] = to;
        balanceOf[from] -= 1;
        balanceOf[to] += 1;
        (bool f, ) = payable(to).call{value: faucet}("");
        require(f);
        emit Transfer(from, to, 1);
    }

    /**
     * @dev Destroys `tokenId`.
     * The approval is cleared when the token is burned.
     *
     * Requirements:
     *
     * - `tokenId` must exist.
     *
     * Emits a {Transfer} event.
     */
    function _burn(uint256 tokenId) internal virtual {
        require(assetsIndexToAddress[tokenId] != address(0));
        address owner = assetsIndexToAddress[tokenId];
        balanceOf[owner] -= 1;
        assetsIndexToAddress[tokenId] = address(0);
        emit Transfer(owner, address(0), tokenId);
    }

    function burn(uint _id) public onlyOwner {
        _burn(_id);
    }

    function withdraw() public onlyOwner{
        // =============================================================================
        // This will transfer the remaining contract balance to the owner.
        // Do not remove this otherwise you will not be able to withdraw the funds.
        // =============================================================================
        payable(msg.sender).transfer(address(this).balance);
        // =============================================================================
    }

    // Function to receive Ether. msg.data must be empty
    receive() external payable {}

    // Fallback function is called when msg.data is not empty
    fallback() external payable {}

}