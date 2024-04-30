import 'package:flutter/material.dart';

class VoyageOverviewDestination extends StatelessWidget {
  final String destinationName;
  final String destinationImage;

  const VoyageOverviewDestination(
      {super.key,
      required this.destinationName,
      required this.destinationImage});

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 100,
      child: Stack(
        alignment: Alignment.center,
        fit: StackFit.expand,
        children: <Widget>[
          Hero(
            tag: destinationName,
            child: Image.network(
              destinationImage,
              fit: BoxFit.cover,
            ),
          ),
          Container(
            color: Colors.black45,
            child: Center(
              child: Text(
                destinationName,
                style: const TextStyle(
                  fontSize: 25,
                  color: Colors.white,
                  fontWeight: FontWeight.bold,
                ),
              ),
            ),
          )
        ],
      ),
    );
  }
}
