import 'package:flutter/material.dart';

class SecondScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return const Scaffold (
      body: Center (
          child: Text(
          "Hello Second Screen",
          style: TextStyle(
          fontSize: 30.0,
          color: Colors.black,
          ),
        ),
      ),
    );
  }
}