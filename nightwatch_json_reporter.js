module.exports = {
  write : function(results, options, done) {
    console.log('NIGHTWATCHJSON\n', JSON.stringify(results));
    done();
  }
};
