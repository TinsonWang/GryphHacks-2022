import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class QrGenerate extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        color:const Color(0xFF001d4a),
        alignment: Alignment.center,
        child: Column (
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Image.asset('../Assets/parkit-logo-transparent.jpg', scale: 2),
            SizedBox(height: 30),
            Container (
              width: MediaQuery.of(context).size.width * 0.7,
              height: MediaQuery.of(context).size.width * 0.5,
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(20),
                color: Colors.white,
              ),
              child: Column (
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Image.asset('../Assets/qr.png'),
                  SizedBox(height: 20),
                  const Text("Scan the QR Code at the machine!"),
                ],
              )
            ),
            Container (
              alignment: Alignment.center,
              width: MediaQuery.of(context).size.width * 0.7,
              height: MediaQuery.of(context).size.width * 0.5,
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(20),
                color: Colors.white,
              )
            ),
            SizedBox(height: 30),
            FlatButton(  
              child: Text('Done', style: TextStyle(fontSize: 20.0),),  
              color: Colors.green,  
              textColor: Colors.white,  
              onPressed: () {},  
            ),
            SizedBox(height: 30),
          ]
        ,)
      )
    );
  }
}