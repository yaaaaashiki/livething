import sys
import json
from http.server import BaseHTTPRequestHandler, HTTPServer

class JsonResponseHandler(BaseHTTPRequestHandler):

    def do_POST(self):
        content_len = int(self.headers.get('content-length'))
        requestBody = self.rfile.read(content_len).decode('UTF-8')
        print('requestBody=' + requestBody)

        jsonData = json.loads(requestBody)

        print('**JSON**')
        print(json.dumps(jsonData, sort_keys=False, indent=4, separators={',',':'}))

        self.send_response(200)
        self.send_header('Content-type','text/json')
        self.end_headers()

server = HTTPServer(('',5000), JsonResponseHandler)
server.serve_forever()