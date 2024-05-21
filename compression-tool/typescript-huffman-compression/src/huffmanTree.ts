// export interface HuffmanNode {
//     freq: number;
// }

// export interface InternalNode extends HuffmanNode {
//     left: HuffmanNode;
//     right: HuffmanNode;
// }

// export interface LeafNode extends HuffmanNode {
//     char: string;
// }

export interface LeafNode {
    type: 'leaf';
    char: string;
    freq: number;
}

export interface InternalNode {
    type: 'internal';
    left: HuffmanNode;
    right: HuffmanNode;
    freq: number;
}

export type HuffmanNode = LeafNode | InternalNode;

export function buildHuffmanTree(frequencies: { [char: string]: number }): HuffmanNode {
    console.log('Building Huffman tree with frequencies:', frequencies);

    const nodes: HuffmanNode[] = Object.entries(frequencies).map(([char, freq]) => {
        const node: LeafNode = { type: 'leaf', char, freq };
        console.log('Created leaf node:', node);
        return node;
    });

    console.log('Initial nodes:', nodes);

    while (nodes.length > 1) {
        nodes.sort((a, b) => {
            if (a.freq === b.freq) {
                if (isLeafNode(a) && isLeafNode(b)) {
                    return a.char.localeCompare(b.char);
                }
                return isLeafNode(a) ? -1 : 1;
            }
            return a.freq - b.freq;
        });

        console.log('Sorted nodes:', nodes);

        const left = nodes.shift() as HuffmanNode;
        const right = nodes.shift() as HuffmanNode;

        console.log('Left node:', left);
        console.log('Right node:', right);

        const internalNode: InternalNode = {
            type: 'internal',
            left,
            right,
            freq: left.freq + right.freq,
        };

        console.log('Created internal node:', internalNode);

        nodes.push(internalNode);

        console.log('Updated nodes:', nodes);
    }

    console.log('Final Huffman tree:', nodes[0]);

    return nodes[0];
}

function isLeafNode(node: HuffmanNode): node is LeafNode {
    return 'char' in node;
}
