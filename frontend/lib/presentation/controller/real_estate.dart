
import 'package:flutter/cupertino.dart';
import 'package:frontend/domain/usecase/real_estate.dart';
import 'package:file_picker/file_picker.dart';

import '../../domain/entity/address.dart';
import '../../domain/entity/pagination.dart';
import '../../domain/entity/real_estate.dart';

class RealEstateController{
  final GetRealEstateUseCase getRealEstateUseCase;
  final UploadImageUseCase uploadImageUseCase;
  final CreateRealEstateUseCase createRealEstateUseCase;


  RealEstateController({required this.getRealEstateUseCase, required this.uploadImageUseCase, required this.createRealEstateUseCase});

  final ValueNotifier<bool> isLoading = ValueNotifier<bool>(false);
  final ValueNotifier<String> isError = ValueNotifier<String>("");
  final ValueNotifier<PaginationEntity<RealEstateEntity>?> realEstate = ValueNotifier<PaginationEntity<RealEstateEntity>?>(null);
  final ValueNotifier<bool> isSuccess = ValueNotifier<bool>(false);
  final ValueNotifier<String> createSuccess = ValueNotifier<String>("");

  Future<void> createRealEstate(
      String street, complement, neighborhood, city, stateAbbr, zipCode, transactionType, imageUrl, value
      ) async{

    AddressEntity address = AddressEntity(
        street: street,
        complement: complement,
        neighborhood: neighborhood,
        city: city,
        stateAbbr: stateAbbr,
        zipCode: zipCode,
    );
    RealEstateEntity realEstate = RealEstateEntity(
      address: address,
      price: double.parse(value),
      transactionType: transactionType,
      imageUrl: imageUrl,
    );

    isLoading.value = true;
    isError.value = "";

    try{
      await createRealEstateUseCase.createRealEstate(realEstate);

      createSuccess.value = "Real estate created successfully";
    } catch(e){
      isError.value = e.toString();

    }finally{
      isLoading.value = false;
      isSuccess.value = true;
    }
  }


  Future<void> getRealEstate(int page, int pageSize) async{

    isLoading.value = true;
    isError.value = "";

    try{
      final res = await getRealEstateUseCase.getRealEstate(page, pageSize);

      realEstate.value = res;
    } catch(e){
      isError.value = e.toString();
    } finally{
      isLoading.value = false;
    }


  }

  Future<String> uploadImage(PlatformFile image) async{
    isLoading.value = true;
    isError.value = "";

    try{
      final res = await uploadImageUseCase.uploadImage(image);

      createSuccess.value = "Image uploaded successfully";
      return res;
    } catch(e){
      isError.value = e.toString();
    } finally{
      isLoading.value = false;
      isSuccess.value = true;
    }
    return "";
  }







}


