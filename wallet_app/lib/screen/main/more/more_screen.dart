import 'package:flutter/material.dart';

class MoreScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Text(
          'This is more screen.',
          style: TextStyle(fontSize: 24),
        ),
      ),
    );
  }
}