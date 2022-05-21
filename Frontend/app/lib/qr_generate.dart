import 'package:flutter/material.dart';

class QrGenerate extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        color:const Color(0xFF001d4a),
        alignment: Alignment.center,
        child: FractionallySizedBox (
          widthFactor: 0.8,
          heightFactor: 0.6,
          alignment: Alignment.center,
          child: ClipRRect(
            borderRadius: BorderRadius.circular(10),
            child: Container (
              color:Colors.white
            )
          ),
        )
      )
    );
  }
}