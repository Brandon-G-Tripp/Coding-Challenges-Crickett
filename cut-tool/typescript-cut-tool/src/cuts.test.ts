import * as fs from 'fs';
import * as path from 'path';
import { cutSecondField } from './cut';

describe('cutSecondField', () => {
    const sampleFile = path.join(__dirname, 'sample.tsv');

    beforeEach(() => {
        const content = 'f0\tf1\tf2\tf3\tf4\tf5\n0\t1\t2\t3\t4\t5\n5\t6\t7\t8\t9\t10\n';
        fs.writeFileSync(sampleFile, content);
    });

    afterEach(() => {
        fs.unlinkSync(sampleFile);
    });

    it('should print the second field from each line', async () => {
        const consoleSpy = jest.spyOn(console, 'log');
        await cutSecondField(sampleFile);
        expect(consoleSpy).toHaveBeenCalledWith('f1');
        expect(consoleSpy).toHaveBeenCalledWith('1');
        expect(consoleSpy).toHaveBeenCalledWith('6');
        consoleSpy.mockRestore();
    });
});
