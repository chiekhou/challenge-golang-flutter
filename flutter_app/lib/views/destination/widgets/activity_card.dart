import 'package:flutter/material.dart';
import '../../../models/activity_model.dart';

class ActivityCard extends StatelessWidget {
  final Activity activity;
  final bool isSelected;
  final VoidCallback toggleActivity;


  ActivityCard({
    required this.activity,
    required this.isSelected,
    required this.toggleActivity,

  });


  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: toggleActivity,
      child: Card(
        child: Stack(
          children: [
            Ink.image(
              image: NetworkImage(activity.image),
              fit: BoxFit.cover,
              child: InkWell(
                onTap:  toggleActivity,
              ),
            ),
            if (isSelected)
              Container(
                color: Colors.black.withOpacity(0.5),
                child: Center(
                  child: Icon(
                    Icons.check,
                    color: Colors.white,
                    size: 48.0,
                  ),
                ),
              ),
            Positioned(
              bottom: 0,
              left: 0,
              right: 0,
              child: Container(
                color: Colors.black.withOpacity(0.7),
                padding: EdgeInsets.all(8.0),
                child: Text(
                  activity.name,
                  style: TextStyle(
                    color: Colors.white,
                    fontSize: 16,
                  ),
                  textAlign: TextAlign.center,
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}