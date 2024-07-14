import 'package:flutter/material.dart';

class BudgetManagementScreen extends StatelessWidget {
  static const routeName = '/admin/budget';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Gestion du budget'),
      ),
      body: Center(
        child: Text(
          'Gestion du budget des groupes',
          style: TextStyle(fontSize: 24),
        ),
      ),
    );
  }
}
