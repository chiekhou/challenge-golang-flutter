import 'package:flutter/material.dart';
import '../../../models/destination_model.dart';

class VoyageDestinationBar extends StatelessWidget {
  final Destination destination;

  const VoyageDestinationBar({super.key, required this.destination});

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 200,
      child: Stack(
        fit: StackFit.expand,
        alignment: Alignment.center,
        children: <Widget>[
          Image.network(
            destination.image,
            fit: BoxFit.cover,
          ),
          Container(
            color: Colors.black38,
            padding: const EdgeInsets.symmetric(
              vertical: 30,
              horizontal: 10,
            ),
            child: Column(
              children: <Widget>[
                Row(
                  children: <Widget>[
                    IconButton(
                      icon: const Icon(
                        Icons.arrow_back,
                        size: 30,
                        color: Colors.white,
                      ),
                      onPressed: () => Navigator.pop(context),
                    ),
                  ],
                ),
                Expanded(
                  child: Center(
                    child: Text(
                      destination.name,
                      style: const TextStyle(
                        fontSize: 30,
                        color: Colors.white,
                      ),
                    ),
                  ),
                )
              ],
            ),
          )
        ],
      ),
    );
  }
}
