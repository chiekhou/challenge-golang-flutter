import 'package:flutter/material.dart';
import 'package:flutter_app/widgets/content.dart';
import 'package:flutter_app/widgets/sidebar.dart';

class DashboardScreen extends StatefulWidget {
  static const routeName = '/admin/dashboard';
  @override
  _DashboardScreenState createState() => _DashboardScreenState();
}

class _DashboardScreenState extends State<DashboardScreen> {
  String _selectedCategory = 'Home';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Row(
        children: [
          Sidebar(
            onCategorySelected: (category) {
              setState(() {
                _selectedCategory = category;
              });
            },
          ),
          Expanded(
            child: ContentArea(
              category: _selectedCategory,
            ),
          ),
        ],
      ),
    );
  }
}
