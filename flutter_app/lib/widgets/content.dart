import 'package:flutter/material.dart';
import 'package:flutter_app/views/admin/group_management_screen.dart';
import 'package:flutter_app/views/admin/home_dashboard.dart';
import 'package:flutter_app/views/admin/user_management_screen.dart';
import 'package:flutter_app/views/admin/voyage_management_screen.dart';

class ContentArea extends StatelessWidget {
  final String category;

  const ContentArea({required this.category});

  @override
  Widget build(BuildContext context) {
    Widget content;
    switch (category) {
      case 'Home':
        content = const DashboardHomeScreen();
        break;
      case 'Users':
        content = const UserManagementScreen();
        break;
      case 'Groups':
        content = const GroupManagementScreen();
        break;
      case 'Trips':
        content = VoyageManagementScreen();
        break;
      default:
        content = const DashboardHomeScreen();
    }

    return Container(
      padding: const EdgeInsets.all(16.0),
      child: content,
    );
  }
}
