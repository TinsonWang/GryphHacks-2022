import 'package:flutter/material.dart';

void main() {
  runApp(MaterialApp(
    home: SplashPage()
  ));
}

class SplashPage extends StatelessWidget {

  @override
  Widget build(BuildContext context) {

    return Scaffold(
      body: Container(
        color:Colors.blueAccent,
        alignment: Alignment.center,
          child: Image.asset('assets/parkit-logo-transparent.jpg')
      )
    );
  }
}