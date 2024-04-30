import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_app/providers/destination_provider.dart';
import 'package:image_picker/image_picker.dart';
import 'package:provider/provider.dart';

class ActivityFormImagePicker extends StatefulWidget {
  final Function updateUrl;

  const ActivityFormImagePicker({super.key, required this.updateUrl});

  @override
  State<ActivityFormImagePicker> createState() =>
      _ActivityFormImagePickerState();
}

class _ActivityFormImagePickerState extends State<ActivityFormImagePicker> {
  File? _deviceImage;
  final picker = ImagePicker();

  Future<void> _pickImage(ImageSource source) async {
    try {
      XFile? pickedFile = await picker.pickImage(source: source);
      if (pickedFile != null && mounted) {
        _deviceImage = File(pickedFile.path);
        final url =
            await Provider.of<DestinationProvider>(context, listen: false)
                .uploadImage(_deviceImage!);
        widget.updateUrl(url);
        setState(() {});
      } else {}
    } catch (e) {
      rethrow;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            TextButton.icon(
              icon: const Icon(Icons.photo),
              label: const Text('Galerie'),
              onPressed: () => _pickImage(ImageSource.gallery),
            ),
            TextButton.icon(
              icon: const Icon(Icons.photo_camera),
              label: const Text('camera'),
              onPressed: () => _pickImage(ImageSource.camera),
            ),
          ],
        ),
        SizedBox(
          width: double.infinity,
          child: _deviceImage != null
              ? Image.file(
                  _deviceImage!,
                  fit: BoxFit.cover,
                )
              : const Text('Aucune image'),
        )
      ],
    );
  }
}
