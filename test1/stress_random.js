import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
    scenarios: {
        contacts: {
            executor: "ramping-vus",
            startVUs: 0,
            stages: [
                { duration: "5m", target: 500 },
                { duration: "5m", target: 1000 },
                { duration: "5m", target: 2000 },
            ],
            gracefulRampDown: "0s",
        },
    },
};

export default function () {
    let response;
    const url = "Your Key";
    const data = [
        {
            jsonrpc: "2.0",
            method: "eth_blockNumber",
            params: [],
            id: 1,
        },
        {
            jsonrpc: "2.0",
            method: "eth_getTransactionReceipt",
            params: [
                "0xabfe871f8bdb0f526611e8620a4569cdd7617f7e2bbc377fe6ab5b2a27e76d6e",
            ],
            id: 1,
        },
        {
            jsonrpc: "2.0",
            method: "eth_getTransactionByBlockNumberAndIndex",
            params: ["latest", "0x0"],
            id: 1,
        },
    ];
    const randMethod = data[Math.floor(Math.random() * data.length)];
    const headers = { headers: { "Content-Type": "application/json" },};

    response = http.post(url, JSON.stringify(randMethod), headers);
    check(response, { 'Is status = 200': (r) => r.status == 200 });
    check(response, { 'Is status != 200': (r) => r.status != 200 });
    check(response, {
      'Is error = hypernode not found': (r) =>
        r.body.includes('hypernode not found'),
    });
    check(response, {
      'Is error =  timeout': (r) =>
        r.body.includes('timeout'),
    });
    //console.log(response);
    sleep(1);
}