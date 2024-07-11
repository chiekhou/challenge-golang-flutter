import 'package:flutter/material.dart';
import 'package:flutter_app/models/hotel_model.dart';

class HotelCard extends StatelessWidget {
  final Hotel hotel;
  final bool isSelected;
  final VoidCallback toggleHotel;

  HotelCard({
    required this.hotel,
    required this.isSelected,
    required this.toggleHotel,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: toggleHotel,
      child: Card(
        child: Stack(
          children: [
            Ink.image(
              image: NetworkImage(hotel.image),
              fit: BoxFit.cover,
              child: InkWell(
                onTap:  toggleHotel,
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
                  hotel.name,
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