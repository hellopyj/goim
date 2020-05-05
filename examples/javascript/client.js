var Ws;
(function(win) {
    const rawHeaderLen = 16;
    const packetOffset = 0;
    const headerOffset = 4;
    const verOffset = 6;
    const opOffset = 8;
    const seqOffset = 12;

    var Client = function(options) {
        var MAX_CONNECT_TIMES = 10;
        var DELAY = 15000;
        this.options = options || {};
        this.createConnect(MAX_CONNECT_TIMES, DELAY);
    }

    var appendMsg = function(text) {
        var span = document.createElement("SPAN");
        var text = document.createTextNode(text);
        span.appendChild(text);
        document.getElementById("box").appendChild(span);
    }

    Client.prototype.createConnect = function(max, delay) {
        var self = this;
        if (max === 0) {
            return;
        }
        connect();

        var textDecoder = new TextDecoder();
        var textEncoder = new TextEncoder();
        var heartbeatInterval;
        function connect() {
            var ws = new WebSocket('ws://127.0.0.1:9326?room=204066851947286529&uid=202331638229635072');
           //var ws = new WebSocket('ws://127.0.0.1:9326/?name=%E7%AA%A6%E6%98%A5%E9%BD%90&tiows_reconnect=true');
            Ws=ws;
            ws.sendMsg=function (msg) {
                var headerBuf = new ArrayBuffer(rawHeaderLen);
                var headerView = new DataView(headerBuf, 0);
                var bodyBuf = textEncoder.encode(msg);
                headerView.setInt32(packetOffset, rawHeaderLen + bodyBuf.byteLength);
                headerView.setInt16(headerOffset, rawHeaderLen);
                headerView.setInt16(verOffset, 1);
                headerView.setInt32(opOffset, 1002);
                headerView.setInt32(seqOffset, 100000);
                ws.send(mergeArrayBuffer(headerBuf, bodyBuf));

            }
            ws.sendC=function(msg){
                var token ='live://1001';
                var headerBuf = new ArrayBuffer(rawHeaderLen);
                var headerView = new DataView(headerBuf, 0);
                var bodyBuf = textEncoder.encode(token);
                headerView.setInt32(packetOffset, rawHeaderLen + bodyBuf.byteLength);
                headerView.setInt16(headerOffset, rawHeaderLen);
                headerView.setInt16(verOffset, 1);
                headerView.setInt32(opOffset, 12);
                headerView.setInt32(seqOffset, 1);
                console.log("jjsjka",headerBuf)
                ws.send(mergeArrayBuffer(headerBuf, bodyBuf));

                appendMsg("send: auth token: " + token);
            }
            ws.binaryType = 'arraybuffer';
            ws.onopen = function() {
                auth();
            }

            ws.onmessage = function(evt) {
                var data = evt.data;
                var dataView = new DataView(data, 0);
                var packetLen = dataView.getInt32(packetOffset);
                var headerLen = dataView.getInt16(headerOffset);
                var ver = dataView.getInt16(verOffset);
                var op = dataView.getInt32(opOffset);
                var seq = dataView.getInt32(seqOffset);

                console.log("receiveHeader: packetLen=" + packetLen, "headerLen=" + headerLen, "ver=" + ver, "op=" + op, "seq=" + seq);

                switch(op) {
                    case 8:
                        // auth reply ok
                        document.getElementById("status").innerHTML = "<color style='color:green'>ok<color>";
                        appendMsg("receive: auth reply");
                        // send a heartbeat to server
                        heartbeat();
                        heartbeatInterval = setInterval(heartbeat, 30 * 1000);
                        break;
                    case 3:
                        // receive a heartbeat from server
                        console.log("receive: heartbeat");
                        appendMsg("receive: heartbeat reply");
                        break;
                    case 9:
                        // batch message
                        for (var offset=rawHeaderLen; offset<data.byteLength; offset+=packetLen) {
                            // parse
                            var packetLen = dataView.getInt32(offset);
                            var headerLen = dataView.getInt16(offset+headerOffset);
                            var ver = dataView.getInt16(offset+verOffset);
                            var op = dataView.getInt32(offset+opOffset);
                            var seq = dataView.getInt32(offset+seqOffset);
                            var msgBody = textDecoder.decode(data.slice(offset+headerLen, offset+packetLen));
                            // callback
                            messageReceived(ver, msgBody);
                            appendMsg("receive: ver=" + ver + " op=" + op + " seq=" + seq + " message=" + msgBody);
                        }
                        break;
                    default:
                        var msgBody = textDecoder.decode(data.slice(headerLen, packetLen));
                        messageReceived(ver, msgBody);
                        appendMsg("receive: ver=" + ver + " op=" + op + " seq=" + seq + " message=" + msgBody);
                        break
                }
            }

            ws.onclose = function() {
                if (heartbeatInterval) clearInterval(heartbeatInterval);
                setTimeout(reConnect, delay);

                document.getElementById("status").innerHTML =  "<color style='color:red'>failed<color>";
            }

            function heartbeat() {
                var headerBuf = new ArrayBuffer(rawHeaderLen);
                var headerView = new DataView(headerBuf, 0);
                headerView.setInt32(packetOffset, rawHeaderLen);
                headerView.setInt16(headerOffset, rawHeaderLen);
                headerView.setInt16(verOffset, 1);
                headerView.setInt32(opOffset, 2);
                headerView.setInt32(seqOffset, 1);
                ws.send(headerBuf);
            }

            function auth(stoken) {
                var token =stoken?stoken: '{"mid":123,"token":"fa4a41ae-d6f7-427e-bd9b-3b66bc2e6b58", "room_id":"live://1000", "platform":"web", "accepts":[1000,1001,1002]}'
                var headerBuf = new ArrayBuffer(rawHeaderLen);
                var headerView = new DataView(headerBuf, 0);
                var bodyBuf = textEncoder.encode(token);
                headerView.setInt32(packetOffset, rawHeaderLen + bodyBuf.byteLength);
                headerView.setInt16(headerOffset, rawHeaderLen);
                headerView.setInt16(verOffset, 1);
                headerView.setInt32(opOffset, 7);
                headerView.setInt32(seqOffset, 1);
                console.log("jjsjka",headerBuf)
                ws.send(mergeArrayBuffer(headerBuf, bodyBuf));

                appendMsg("send: auth token: " + token);
            }

            function messageReceived(ver, body) {
                var notify = self.options.notify;
                if(notify) notify(body);
                console.log("messageReceived:", "ver=" + ver, "body=" + body);
            }

            function mergeArrayBuffer(ab1, ab2) {
                var u81 = new Uint8Array(ab1),
                    u82 = new Uint8Array(ab2),
                    res = new Uint8Array(ab1.byteLength + ab2.byteLength);
                console.log("sjjs",u81)
                res.set(u81, 0);
                res.set(u82, ab1.byteLength);
                return res.buffer;
            }

            function char2ab(str) {
                var buf = new ArrayBuffer(str.length);
                var bufView = new Uint8Array(buf);
                for (var i=0; i<str.length; i++) {
                    bufView[i] = str[i];
                }
                return buf;
            }

        }

        function reConnect() {
            self.createConnect(--max, delay * 2);
        }
    }

    win['MyClient'] = Client;

})(window);

