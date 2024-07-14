import 'package:flutter/material.dart';
import 'user_management_screen.dart';
import 'group_management_screen.dart';
import 'voyage_management_screen.dart';
import 'destination_management_screen.dart';
import 'activity_management_screen.dart';
import 'budget_management_screen.dart';

class AdminDashboardScreen extends StatelessWidget {
  static const routeName = '/admin/dashboard';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Tableau de bord'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            ElevatedButton(
              onPressed: () {
                Navigator.pushNamed(context, UserManagementScreen.routeName);
              },
              child: Text('Gestion des utilisateurs'),
            ),
            ElevatedButton(
              onPressed: () {
                Navigator.pushNamed(context, GroupManagementScreen.routeName);
              },
              child: Text('Gestion des groupes'),
            ),
            ElevatedButton(
              onPressed: () {
                Navigator.pushNamed(context, VoyageManagementScreen.routeName);
              },
              child: Text('Gestion des voyages'),
            ),
            ElevatedButton(
              onPressed: () {
                Navigator.pushNamed(
                    context, DestinationManagementScreen.routeName);
              },
              child: Text('Gestion des destinations'),
            ),
            ElevatedButton(
              onPressed: () {
                Navigator.pushNamed(
                    context, ActivityManagementScreen.routeName);
              },
              child: Text('Gestion des activités'),
            ),
            ElevatedButton(
              onPressed: () {
                Navigator.pushNamed(context, BudgetManagementScreen.routeName);
              },
              child: Text('Gestion du budget'),
            ),
          ],
        ),
      ),
    );
  }
}
