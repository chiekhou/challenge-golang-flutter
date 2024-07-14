import 'package:flutter/material.dart';

class UserManagementScreen extends StatelessWidget {
  static const routeName = '/admin/users';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Gestion des utilisateurs'),
      ),
      body: Center(
        child: Text(
          'Liste des utilisateurs',
          style: TextStyle(fontSize: 24),
        ),
      ),
    );
  }
}
