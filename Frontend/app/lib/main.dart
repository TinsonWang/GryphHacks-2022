import 'package:flutter/material.dart';
import 'splash_screen.dart';
import 'qr_generate.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: SplashScreen(),
    );
  }
}
