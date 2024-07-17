import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../../../providers/voyage_provider.dart';

class VoyageDropdownMenu extends StatefulWidget {
  final Function(int?) onSelectedVoyageChanged;

  const VoyageDropdownMenu({Key? key, required this.onSelectedVoyageChanged}) : super(key: key);

  @override
  _VoyageDropdownMenuState createState() => _VoyageDropdownMenuState();
}

class _VoyageDropdownMenuState extends State<VoyageDropdownMenu> {
  int? _selectedVoyageId;

  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addPostFrameCallback((_) {
      Provider.of<VoyageProvider>(context, listen: false).fetchData();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Consumer<VoyageProvider>(
      builder: (context, voyageProvider, child) {
        if (voyageProvider.isLoading) {
          return CircularProgressIndicator();
        } else if (voyageProvider.voyages.isEmpty) {
          return Text('Aucun voyage disponible');
        } else {
          return DropdownButton<int>(
            value: _selectedVoyageId,
            hint: Text('Voulez-vous lier votre groupe à un voyage?'),
            items: voyageProvider.voyages.map((voyage) {
              return DropdownMenuItem<int>(
                value: voyage.id,
                child: Text(voyage.destination),
              );
            }).toList(),
            onChanged: (int? newValue) {
              setState(() {
                _selectedVoyageId = newValue;
              });
              widget.onSelectedVoyageChanged(newValue);
            },
          );
        }
      },
    );
  }
}
