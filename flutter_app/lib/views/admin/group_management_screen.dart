import 'package:flutter/material.dart';

class GroupManagementScreen extends StatelessWidget {
  static const routeName = '/admin/groups';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Gestion des groupes'),
      ),
      body: Center(
        child: Text(
          'Liste des groupes',
          style: TextStyle(fontSize: 24),
        ),
      ),
    );
  }
}
