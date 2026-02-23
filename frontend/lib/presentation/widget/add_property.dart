import 'dart:io';
import 'dart:math';

import 'package:flutter/cupertino.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:frontend/domain/entity/address.dart';
import 'package:image_picker/image_picker.dart';

import '../../l10n/app_localizations.dart';
import '../controller/address.dart';
import '../controller/real_estate.dart';

class AddProperty extends StatefulWidget {
  final RealEstateController realEstateController;
  final AddressController addressController;

  final Function(String page) navigate;


  const AddProperty({super.key, required this.navigate, required this.realEstateController, required this.addressController});

  @override
  State<AddProperty> createState() => _AddPropertyState();
}

class _AddPropertyState extends State<AddProperty> {
  File? _pickedImage;
  String imageUrl = "";
  AddressEntity? address;
  final List<String> _apiValues = ['RENT', 'SELL'];
  String _currentTechnicalValue = 'RENT';


  Future<void> getImage() async {
    final ImagePicker _picker = ImagePicker();
    final XFile? image = await _picker.pickImage(source: ImageSource.gallery);

    if (image == null) return;
    setState(() {
      _pickedImage = File(image.path);
    });

    imageUrl = await widget.realEstateController.uploadImage(image);
    return;
  }

  final TextEditingController _inputZipCode = TextEditingController(text: '');

  final TextEditingController _inputStreet = TextEditingController(text: '');
  final TextEditingController _inputNeighborhood = TextEditingController(text: '');
  final TextEditingController _inputComplement = TextEditingController(text: '');
  final TextEditingController _inputCity = TextEditingController(text: '');
  final TextEditingController _inputState = TextEditingController(text: '');
  final TextEditingController _inputValue = TextEditingController(text: '');





  Future<void> getAddress() async{
    String zipCode = _inputZipCode.text;


    AddressEntity? res = await widget.addressController.getAddress(zipCode);

    res ??= AddressEntity(zipCode: "", street: "", complement: "", neighborhood: "", city: "", stateAbbr: "");
      setState(() {
        address = res;
        res?.street != "" ?
          _inputStreet.text = "${res?.street}, ${res?.complement}" :
            _inputStreet.text = "";
        _inputNeighborhood.text = res!.neighborhood;
        _inputCity.text = res.city;
        _inputState.text = res.stateAbbr;
      });
    }


  Future<void> createRealEstate() async{

    await widget.realEstateController.createRealEstate(
        _inputStreet.text, _inputComplement.text, _inputNeighborhood.text, _inputCity.text, _inputState.text, _inputZipCode.text, _currentTechnicalValue, imageUrl, _inputValue.text
    );

    setState(() {
      address = null;
      _inputZipCode.text = "";
      _pickedImage = null;
      imageUrl = "";
    });


  }



  @override
  Widget build(BuildContext context) {
    const primaryDarkColor = Color(0xFF131A2A);
    const borderColor = Color(0xFFE5E7EB);
    final l10n = AppLocalizations.of(context)!;

    final Map<String, String> translations = {
      'RENT': l10n.rent,
      'SELL': l10n.sell,
    };

    return SingleChildScrollView(
      child: Center(
      child: ConstrainedBox(
        constraints: const BoxConstraints(
          maxWidth: 600,
        ),
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Container(
                decoration: BoxDecoration(
                  color: const Color(0xFFF3F4F6),
                  borderRadius: BorderRadius.circular(8),
                ),
                child: TextButton.icon(
                  onPressed: () => widget.navigate("home"),
                  icon: const Icon(Icons.arrow_back, size: 18, color: primaryDarkColor),
                  label: Text(
                    l10n.backButton,
                    style: TextStyle(
                      color: primaryDarkColor,
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                  style: TextButton.styleFrom(
                    padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                ),
              ),
              const SizedBox(height: 16),
      
              Container(
                padding: const EdgeInsets.all(24),
                decoration: BoxDecoration(
                  color: Colors.white,
                  border: Border.all(color: borderColor),
                  borderRadius: BorderRadius.circular(12),
                  boxShadow: [
                    BoxShadow(
                      color: Colors.black.withOpacity(0.02),
                      blurRadius: 8,
                      offset: const Offset(0, 4),
                    ),
                  ],
                ),
                child: Column(
                    children: [
                      GestureDetector(
                        onTap: () => getImage(),
                        child: Container(
                          width: double.infinity,
                          height: 200,
                          decoration: BoxDecoration(
                            color: Colors.grey[200],
                            borderRadius: BorderRadius.circular(12),
                            border: Border.all(color: Colors.grey, style: BorderStyle.solid),
                          ),
                          child: _pickedImage != null
                              ? ClipRRect(
                            borderRadius: BorderRadius.circular(12),
                            child: kIsWeb
                          ? Image.network(_pickedImage!.path, fit: BoxFit.cover)
                              : Image.file(File(_pickedImage!.path), fit: BoxFit.cover)
                          )
                              :  Column(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              Icon(Icons.add_a_photo, size: 50, color: Colors.grey),
                              SizedBox(height: 8),
                              Text(l10n.addImage, style: TextStyle(color: Colors.grey)),
                            ],
                          ),
                        ),
                      ),
      
                      const SizedBox(height: 24),

                      Row(
                        children: [
                          Expanded(
                            child: TextField(
                              controller: _inputZipCode,
                              keyboardType: const TextInputType.numberWithOptions(decimal: false),
                              decoration: InputDecoration(
                                labelText: l10n.zipCode,
                                border: const OutlineInputBorder(),
                              ),
                            ),
                          ),
                          const SizedBox(width: 8),
                          ElevatedButton(
                            style: ElevatedButton.styleFrom(
                              padding: const EdgeInsets.symmetric(vertical: 16, horizontal: 16),
                            ),
                            onPressed: () => getAddress(),
                            child: Text(l10n.searchZipCode, style: TextStyle(fontSize: 16)),
                          ),
                        ],
                      ),
                      const SizedBox(height: 16),

                      if (address != null)
                        Column(
                        children: [
                          TextField(
                          controller: _inputStreet,
                          keyboardType: TextInputType.text,
                          decoration: InputDecoration(
                            labelText: l10n.street,
                            border: OutlineInputBorder(),
                          ),
                        ),


                          const SizedBox(height: 16),
                          Row(
                            children: [

                              Expanded(
                                child: TextField(
                                  controller: _inputComplement,
                                  keyboardType: TextInputType.text,
                                  decoration: InputDecoration(
                                    labelText: l10n.complement,
                                    border: OutlineInputBorder(),
                                  ),
                                ),
                              ),
                              Expanded(
                                child: TextField(
                                  controller: _inputNeighborhood,
                                  keyboardType: TextInputType.text,
                                  decoration: InputDecoration(
                                    labelText: l10n.neighborhood,
                                    border: OutlineInputBorder(),
                                  ),
                                ),
                              ),
                            ],
                          ),

                          const SizedBox(height: 16),

                          Row(
                            children: [
                              Expanded(
                                child: TextField(
                                  controller: _inputCity,
                                  keyboardType: TextInputType.text,
                                  decoration: InputDecoration(
                                    labelText: l10n.city,
                                    border: OutlineInputBorder(),
                                  ),
                                ),
                              ),
                              Expanded(
                                child: TextField(
                                  controller: _inputState,
                                  keyboardType: TextInputType.text,
                                  decoration: InputDecoration(
                                    labelText: l10n.state,
                                    border: OutlineInputBorder(),
                                  ),
                                ),
                              ),
                            ],
                          ),

                          const SizedBox(height: 16),

                          TextField(
                            controller: _inputValue,
                            keyboardType: const TextInputType.numberWithOptions(decimal: true),
                            decoration: InputDecoration(
                              labelText: l10n.value,
                              border: OutlineInputBorder(),
                            ),
                          ),

                          const SizedBox(height: 16),

                          InputDecorator(
                            decoration: InputDecoration(
                              labelText: l10n.contractType,
                              contentPadding: const EdgeInsets.symmetric(horizontal: 12, vertical: 15),
                              border: OutlineInputBorder(
                                borderRadius: BorderRadius.circular(8),
                              ),
                            ),
                            child: DropdownButtonHideUnderline(
                              child: DropdownButton<String>(
                                value: _currentTechnicalValue,
                                items: _apiValues.map((String value) {
                                  return DropdownMenuItem<String>(
                                    value: value,
                                    child: Text(translations[value] ?? value),
                                  );
                                }).toList(),
                                onChanged: (String? newValue) {
                                  if (newValue != null) {
                                    setState(() {
                                      _currentTechnicalValue = newValue;
                                    });
                                  }
                                },
                              ),
                            ),
                          )





                        ],
                      ),

                      const SizedBox(height: 16),

                      if (address != null)
                        SizedBox(
                          width: double.infinity,
                          child: ElevatedButton(
                            style: ElevatedButton.styleFrom(
                              padding: const EdgeInsets.symmetric(vertical: 16),
                            ),
                            onPressed: createRealEstate,
                            child: Text(l10n.saveProperty, style: TextStyle(fontSize: 16)),
                          ),
                        ),
                    ],
                  ),
              ),
            ],
          ),
        ),
      ),
    ),
  );
  }
}
