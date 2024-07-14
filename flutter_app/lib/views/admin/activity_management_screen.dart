import 'package:flutter/material.dart';

class ActivityManagementScreen extends StatelessWidget {
  static const routeName = '/admin/activities';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Gestion des activités'),
      ),
      body: Center(
        child: Text(
          'Liste des activités',
          style: TextStyle(fontSize: 24),
        ),
      ),
    );
  }
}
