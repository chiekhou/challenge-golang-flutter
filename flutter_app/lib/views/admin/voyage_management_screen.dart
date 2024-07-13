import 'package:flutter/material.dart';

class VoyageManagementScreen extends StatelessWidget {
  static const routeName = '/admin/voyages';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Gestion des voyages'),
      ),
      body: Center(
        child: Text(
          'Liste des voyages',
          style: TextStyle(fontSize: 24),
        ),
      ),
    );
  }
}
