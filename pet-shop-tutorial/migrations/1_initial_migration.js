var Migrations = artifacts.require("./token.sol");

module.exports = function(deployer) {
  deployer.deploy(Migrations);
};
