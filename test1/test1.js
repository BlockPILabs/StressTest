import http from 'k6/http';
import { check } from 'k6';


var ip = 'k6.testnet.blockpi.net';
// var key = getkeys();
var key = "";
var url = 'https://' + ip + '/v1/rpc/' + key;

var params = {
    headers: {
        'Content-Type': 'application/json',
    },
};


export default function () {
    var number = 98658835;
    var hash = "";

    var i = Math.floor(Math.random() * 3)
    switch (i) {
        case 0:
            getBlockByNumber(number)
        case 1:
            getTransactionByHash(hash)
        case 2:
            ethGetTransactionReceipt(hash)
    }

//     getBalance()
//     let number = blockNumber();
//     if (number.length > 0) {
//         let hash = getBlockByNumber(number);
//         if (hash.length > 0) {
//             getTransactionByHash(hash);
//         }
//     }
}


function randomAddr() {
    var len = 40;
    var chars = 'abcdef0123456789';
    var maxPos = chars.length;
    var str = '';
    for (var i = 0; i < len; i++) {
        str += chars.charAt(Math.floor(Math.random() * maxPos));
    }
    return '0x' + str;
}

function getBalance() {
    var addr = randomAddr();

    const payload = JSON.stringify({
        "jsonrpc": "2.0",
        "method": "eth_getBalance",
        "params": [
            addr,
            "latest"
        ],
        "id": 1
    });

    http.post(url, payload, params);
}

function blockNumber() {
    const payload = JSON.stringify({
        "jsonrpc": "2.0",
        "method": "eth_blockNumber",
        "params": [],
        "id": 1
    });

    const res = http.post(url, payload, params, );
    // console.info(res)
    check(res, {
        'eth_blockNumber is status 200': (r) => r.status === 200,
    });
    if (res.status != 200) {
        return ""
    }
    return res.json()['result']
}


function getBlockByNumber(number) {
    const payload = JSON.stringify({
        "jsonrpc": "2.0",
        "method": "eth_getBlockByNumber",
        "params": [number, true],
        "id": 1
    });

    const res = http.post(url, payload, params);
    check(res, {
        'eth_getBlockByNumber is status 200': (r) => r.status === 200,
    });

    let hash = ''
    if (res.status === 200) {
        const arr = res.json()['result']
        if (arr != null && "transactions" in arr) {
            if (arr['transactions'].length > 0) {
                hash = arr['transactions'][0]['hash']
            }
        }
    }

    return hash
}


function getTransactionByHash(hash) {
    const payload = JSON.stringify({
        "jsonrpc": "2.0",
        "method": "eth_getTransactionByHash",
        "params": [hash],
        "id": 1
    });

    const res = http.post(url, payload, params);
    check(res, {
        'eth_getTransactionByHash is status 200': (r) => r.status === 200,
    });
}


function ethGetTransactionReceipt(hash) {
    const payload = JSON.stringify({
        "jsonrpc": "2.0",
        "method": "eth_getTransactionReceipt",
        "params": [hash],
        "id": 1
    });

    const res = http.post(url, payload, params);
    check(res, {
        'eth_getTransactionReceipt is status 200': (r) => r.status === 200,
    });
}

function getkeys() {
    var keys = []
    var index = Math.floor((Math.random() * keys.length));
    return keys[index]
}