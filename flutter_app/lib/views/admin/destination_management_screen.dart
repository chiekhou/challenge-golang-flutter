import 'package:flutter/material.dart';

class DestinationManagementScreen extends StatelessWidget {
  static const routeName = '/admin/destinations';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Gestion des destinations'),
      ),
      body: Center(
        child: Text(
          'Liste des destinations',
          style: TextStyle(fontSize: 24),
        ),
      ),
    );
  }
}
