import http from 'k6/http';
import { sleep } from 'k6';

const API_KEY = 'YOUR_API_KEY';  /*REPLACE WITH YOUR KEY*/

export const options = {
    stages: [
        { duration: '5m', target: 500 },
        { duration: '5m', target: 1000 },
        { duration: '5m', target: 2000 },
    ],
};


export default function () {
    const i = Math.floor(Math.random() * 3)
    const requests = [
        {
            jsonrpc: "2.0",
            method: "klay_blockNumber",
            params: [],
            id: 83,
        },
        {
            jsonrpc: "2.0",
            method: "klay_getTransactionReceipt",
            params: [
                "0x3e7bd142101afb263b4af8175bf97095ea2b06d41ab5dec399c0f49076ab22e0",
            ],
            id: 1,
        },
        {
            jsonrpc: "2.0",
            method: "klay_getTransactionByHash",
            params: ["0x3e7bd142101afb263b4af8175bf97095ea2b06d41ab5dec399c0f49076ab22e0"],
            id: 1,
        },
    ];
    const url = 'https://klaytn.testnet.blockpi.net/v1/rpc/' + API_KEY;
    const headers = {
        headers: { "Content-Type": "application/json" },
    };
    const randomRequest = requests[i];
    const response = http.post(url, JSON.stringify(randomRequest), headers);
    console.log(response.json());
    sleep(1);
}
