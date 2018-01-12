/*
  Web client

 This sketch connects to a website (http://www.google.com)
 using an Arduino Wiznet Ethernet shield.

 Circuit:
 * Ethernet shield attached to pins 10, 11, 12, 13

 created 18 Dec 2009
 by David A. Mellis
 modified 9 Apr 2012
 by Tom Igoe, based on work by Adrian McEwen

 */

#include <SPI.h>
#include <Ethernet.h>

int val = 0;

// Enter a MAC address for your controller below.
// Newer Ethernet shields have a MAC address printed on a sticker on the shield
byte mac[] = { 0x90, 0xA2, 0xDA, 0x0F, 0x3F, 0xF4 };
byte ip[] = {133, 2, 208, 174};
byte gateway[] = { 133, 2, 208, 254};
byte subnetmask[] = {255, 255, 255, 128};

// byte server[] = { 133, 2, 208, 182 };
IPAddress server(133,2,208,182);


// Set the static IP address to use if the DHCP fails to assign

// Initialize the Ethernet client library
// with the IP address and port of the server
// that you want to connect to (port 80 is default for HTTP):
EthernetClient client;

void setup() {
  // Open serial communications and wait for port to open:
  Serial.begin(9600);
  while (!Serial) {
    ; // wait for serial port to connect. Needed for native USB port only
  }

  Ethernet.begin(mac, ip, gateway, gateway, subnetmask);
  
  Serial.print("IP Address: ");
  Serial.println(Ethernet.localIP());
  Serial.print("Subnet Mask: ");
  Serial.println(Ethernet.subnetMask());
  Serial.print("Gateway IP Address: "); 
  Serial.println(Ethernet.gatewayIP());
  Serial.print("DNS Server Address: "); 
  Serial.println(Ethernet.dnsServerIP());
  
  // give the Ethernet shield a second to initialize:
  delay(1000);
}

void loop() {
  // if there are incoming bytes available
  // from the server, read them and print them:
  while(1){
    val = analogRead(0);

    if (client.available()) {
      char c = client.read();
      Serial.print(c);
    }else{
      Do_GET();
    }
  
    // if the server's disconnected, stop the client:
    if (!client.connected()) {
      Serial.println();
      Serial.println("disconnecting.");
      client.stop();
    }
  }
}

void Do_GET(){
  Serial.println("connecting...");
  if(client.connect(server, 5000)) {
    Serial.println("connected");

    // if you get a connection, report back via serial:
    // Make a HTTP request:
    client.print("GET /test_server.py?value=");
    client.print(val);
    client.println(" HTTP/1.1");
    client.println("Host: 192.168.71.4");
    //client.println("Connection: close");
    client.println();
  }else{
    // if you didn't get a connection to the server:
    Serial.println("connection failed");
  }  

  delay(1000);
}

