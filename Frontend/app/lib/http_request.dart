import 'package:flutter/material.dart';


/*
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
*/

import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

Future<Album> createAlbum(String login, String password) async {
  final response = await http.post(
    Uri.parse('https://f612-184-144-100-19.ngrok.io/login'),
    body: jsonEncode(<String, String>{
      'login': login,
      'password' : password
    }),
  );

  if (response.statusCode == 201) {
    // If the server did return a 201 CREATED response,
    // then parse the JSON.
    return Album.fromJson(jsonDecode(response.body));
  } else {
    // If the server did not return a 201 CREATED response,
    // then throw an exception.
    throw Exception('Failed to create album.');
  }
}

class Album {
  //final int id;
  final String login_response;

  //const Album({required this.id, required this.title});

  const Album({required this.login_response});

  factory Album.fromJson( Map<String, dynamic> json ) {
    return Album(
        login_response: json['Status']
    );
  }
}


class SecondScreen extends StatefulWidget {
  const SecondScreen({super.key});

  @override
  _MyAppState createState() {
    return _MyAppState();
  }
}

class _MyAppState extends State<SecondScreen> {
  final TextEditingController _controller = TextEditingController();
  final TextEditingController _controller2 = TextEditingController();

  Future<Album>? _futureAlbum;

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Create Data Example',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: Scaffold(
        body: Container(
          alignment: Alignment.center,
          padding: const EdgeInsets.all(8.0),
          child: (_futureAlbum == null) ? buildColumn() : buildFutureBuilder(),
        ),
      ),
    );
  }

  Column buildColumn() {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      children: <Widget>[

        SizedBox( height: 100 ),

        const Text(
            'Login to your\nAccount',
            textAlign: TextAlign.left,
            style: TextStyle( fontSize: 60, color: Color(0xFF001d4a) )
        ),

        SizedBox( height: 50 ),

        SizedBox(
          width: 400,
          child:
          TextField(
            controller: _controller,
            decoration: const InputDecoration(
              labelText: 'Login',
              border: OutlineInputBorder(),
            ),
          ),
        ),

        SizedBox( height: 50 ),

        SizedBox(
          width: 400,
          child:
          TextField(
            controller: _controller2,
            decoration: const InputDecoration(
                labelText: 'Password',
                border: OutlineInputBorder()
            ),
          ),
        ),

        SizedBox( height: 50 ),

        SizedBox(
            width: 400,
            height: 50,
            child:
            ElevatedButton(
              style: ElevatedButton.styleFrom(
                // Foreground color
                onPrimary: Theme.of(context).colorScheme.onPrimary,
                // Background color
                primary: Color(0xFF006992),
              ).copyWith(elevation: ButtonStyleButton.allOrNull(0.0)),

              onPressed: () {
                setState(() {
                  _futureAlbum = createAlbum(
                      _controller.text,
                      _controller2.text
                  );
                });
              },
              child: const Text('Sign in'),
            )
        ),
      ],
    );
  }

  FutureBuilder<Album> buildFutureBuilder() {
    return FutureBuilder<Album>(
      future: _futureAlbum,
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          return Text(snapshot.data!.login_response);
        } else if (snapshot.hasError) {
          return Text('${snapshot.error}');
        }

        return const CircularProgressIndicator();
      },
    );
  }
}