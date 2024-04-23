const fs = require('fs');

function countBytes(filePath) {
    try {
        const data = fs.readFileSync(filePath);
        return data.length;
    } catch (error) {
        if (error.code === 'ENOENT') {
            return null;
        } 
        throw error;
    } 
} 

module.exports = {
    countBytes,
}; 
