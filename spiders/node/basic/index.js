const {saveItem} = require('crawlab-sdk');

(async () => {
    for (let i = 0; i < 100; i++) {
        await saveItem({'hello': 'world'});
    }
})();
