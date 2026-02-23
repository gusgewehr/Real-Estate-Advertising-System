import 'package:frontend/domain/entity/real_estate.dart';
import 'package:image_picker/image_picker.dart';

import '../entity/pagination.dart';

abstract class RealEstateRepository {
  Future<void> createRealEstate(RealEstateEntity realEstate);
  Future<PaginationEntity<RealEstateEntity>> getRealEstate(int page, int pageSize);
  Future<String> uploadImage(XFile image);
}

